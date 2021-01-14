# MiraiEx DCA

Dette er et verktøy for å jevnlig kjøpe Bitcoin fra MiraiEx over tid
ved å bruke metoden DCA / Dollar Cost Averaging hvor man sakte tar en posisjon i markedet,
noe som kan være en fordel i et så volatilt market som Bitcoin.

Programmet kjører lokalt på din egen datamaskin slik at du slipper å stole
på at noen passer godt nok på dine MiraiEx-API-nøkler.

## Bruk

1. Git clone
2. Endre på `secret_example.go`, legg til dine API-nøkler.
3. `go build`
   
## Bruk (Windows)

1. `./miraiex-dca` for å kjøpe en gitt mengde BTC.

Eksempel: 
```
PS miraiex-dca> go run .
main.BalancesResponse{
    main.Balance{Currency:"BTC", Balance:0.00044521, Hold:0, Available:0.00044521},
    main.Balance{Currency:"DAI", Balance:0, Hold:0, Available:0},
    main.Balance{Currency:"ETH", Balance:0.00415867, Hold:0, Available:0.00415867},
    main.Balance{Currency:"LTC", Balance:0, Hold:0, Available:0},
    main.Balance{Currency:"NOK", Balance:896.76, Hold:0, Available:896.76},
    main.Balance{Currency:"XRP", Balance:0, Hold:0, Available:0}
}
2021/01/14 21:25:58 GET https://api.miraiex.com/v2/markets/BTCNOK/ticker
2021/01/14 21:25:58 Response body: {"bid":"334803.0600000000000000","ask":"342523.0400000000000000","spread":"7719.9800000000000000"}
2021/01/14 21:25:58 Current BTC price: {BTCNOK 334803.06 342523.04 7719.98}
2021/01/14 21:25:58 Request body{"market":"BTCNOK","type":"Bid","price":"342523.04","amount":"0.000100"}
2021/01/14 21:25:59 Response body:{"id":12345678}
2021/01/14 21:25:59 Created order 12345678 for 0.100000 mBTC @ 342523.04 NOK, for a total price of 34.25 NOK + 0.5% fees
```

## Bruk (Linux)

1. Lag en systemd cronjob som regelmessig kjører programmet.