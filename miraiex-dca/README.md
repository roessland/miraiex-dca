# MiraiEx DCA

Command Line Interface (CLI) for å kjøpte BitCoin på MiraiEx.


## TOODO

2021/01/27 23:12:28 Request body: {"market":"BTCNOK","type":"Bid","price":"269861.54","amount":"0.000100"}
2021/01/27 23:12:28 Response body:{"name":"SecurityLevelTooLow","message":"The user's security level is too low"}
2021/01/27 23:12:28 Created order  for 0.100000 mBTC @ 269861.54 NOK, for a total price of 26.99 NOK + 0.5% fees

## miraiex-dca cron

Kommandoen `miraiex-dca cron` kjøres som en scheduled jobb hvert femte minutt,
og utfører alle kjøp og andre bakgrunnsjobber. (TODO)

## miraiex-dca web

Kommandoen `miraiex-dca web` starter en lokal webserver hvor du kan vise statistikk over dine kjøp,
og beregne gevinst/tap per transaksjon. (TODO)

## Epost

Hvis du konfigurerer Mailgun i config-filen kan du få oppdateringer på epost.
F.eks.
* Du har for lite NOK hos MiraiEx -- fyll på (TODO)
* Du kjøpte så og så mye Bitcoin (TODO)
* Månedlige oppdateringer på kjøp/gevinst/tap. (TODO)
* Du har for høye verdier på MiraiEx -- flytt til egen wallet. (TODO)