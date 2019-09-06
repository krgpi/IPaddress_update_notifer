# IPaddress_update_notifer
マシンのグローバルIPアドレスをSlackに返します。

## background 
何らかの事情で、リモートのグローバルIPがコロコロ変わる状況において、IPアドレスをSlack経由で教えてくれるSlackBotです。

デフォルトでは、botがいるチャンネルで`ipaddr`と打つと、IPアドレスが返ってきます。

また`checkUpdate()`を定期実行するように改良することで、IPアドレスが変わった時に、自動でSlackに通知してくれるようになります。

# instllation
`slackRTMconn.go`の`postURL = "https://hooks.slack.com/***"`の部分にはWeb hookのURLを、,`Slack.New("xoxb-***")`にはRTM APIのアクセスキーを書いてください。
