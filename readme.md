# 加密传输的Cosmos-sdk chain（created by starport）
**planet** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).
## 主要工作
1. 在packet传输中使用base64加密Content字段
2. 在ibcpost.go中OnRecvIbcPostPacket解密Content字段
## 效果展示
加密前
```
{"height":"299","txhash":"DB48749AF4F7B93F576F6E14925EBBFB0B163EFC2E9247828DAF8010000750CE","codespace":"","code":0,"data":"0A0D0A0B53656E64496263506F7374","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"SendIbcPost\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]},{\"type\":\"send_packet\",\"attributes\":[{\"key\":\"packet_data\",\"value\":\"\\u0012X\\n\\u0005Hello\\u0012 Hello Mars, I'm Alice from Earth\\u001a-cosmos18tstt6lmjlqdy50xn4c9kasrmjy2l7eyum47pw\"},{\"key\":\"packet_timeout_height\",\"value\":\"0-0\"},{\"key\":\"packet_timeout_timestamp\",\"value\":\"1631715092032909294\"},{\"key\":\"packet_sequence\",\"value\":\"1\"},{\"key\":\"packet_src_port\",\"value\":\"blog\"},{\"key\":\"packet_src_channel\",\"value\":\"channel-0\"},{\"key\":\"packet_dst_port\",\"value\":\"blog\"},{\"key\":\"packet_dst_channel\",\"value\":\"channel-0\"},{\"key\":\"packet_channel_ordering\",\"value\":\"ORDER_UNORDERED\"},{\"key\":\"packet_connection\",\"value\":\"connection-0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"SendIbcPost"},{"key":"module","value":"ibc_channel"}]},{"type":"send_packet","attributes":[{"key":"packet_data","value":"\u0012X\n\u0005Hello\u0012 Hello Mars, I'm Alice from Earth\u001a-cosmos18tstt6lmjlqdy50xn4c9kasrmjy2l7eyum47pw"},{"key":"packet_timeout_height","value":"0-0"},{"key":"packet_timeout_timestamp","value":"1631715092032909294"},{"key":"packet_sequence","value":"1"},{"key":"packet_src_port","value":"blog"},{"key":"packet_src_channel","value":"channel-0"},{"key":"packet_dst_port","value":"blog"},{"key":"packet_dst_channel","value":"channel-0"},{"key":"packet_channel_ordering","value":"ORDER_UNORDERED"},{"key":"packet_connection","value":"connection-0"}]}]}],"info":"","gas_wanted":"200000","gas_used":"54362","tx":null,"timestamp":""}
```
加密后
```
{"height":"853","txhash":"948EF9691CA1884A905837A404C3FD994434BE42C6239C9EADFC0C5F4F6E709E","codespace":"","code":0,"data":"0A0D0A0B53656E64496263506F7374","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"SendIbcPost\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]},{\"type\":\"send_packet\",\"attributes\":[{\"key\":\"packet_data\",\"value\":\"\\u0012d\\n\\u0005Hello\\u0012,SGVsbG8gTWFycywgSSdtIEFsaWNlIGZyb20gRWFydGg=\\u001a-cosmos18tstt6lmjlqdy50xn4c9kasrmjy2l7eyum47pw\"},{\"key\":\"packet_timeout_height\",\"value\":\"0-0\"},{\"key\":\"packet_timeout_timestamp\",\"value\":\"1631715218179045524\"},{\"key\":\"packet_sequence\",\"value\":\"2\"},{\"key\":\"packet_src_port\",\"value\":\"blog\"},{\"key\":\"packet_src_channel\",\"value\":\"channel-0\"},{\"key\":\"packet_dst_port\",\"value\":\"blog\"},{\"key\":\"packet_dst_channel\",\"value\":\"channel-0\"},{\"key\":\"packet_channel_ordering\",\"value\":\"ORDER_UNORDERED\"},{\"key\":\"packet_connection\",\"value\":\"connection-0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"SendIbcPost"},{"key":"module","value":"ibc_channel"}]},{"type":"send_packet","attributes":[{"key":"packet_data","value":"\u0012d\n\u0005Hello\u0012,SGVsbG8gTWFycywgSSdtIEFsaWNlIGZyb20gRWFydGg=\u001a-cosmos18tstt6lmjlqdy50xn4c9kasrmjy2l7eyum47pw"},{"key":"packet_timeout_height","value":"0-0"},{"key":"packet_timeout_timestamp","value":"1631715218179045524"},{"key":"packet_sequence","value":"2"},{"key":"packet_src_port","value":"blog"},{"key":"packet_src_channel","value":"channel-0"},{"key":"packet_dst_port","value":"blog"},{"key":"packet_dst_channel","value":"channel-0"},{"key":"packet_channel_ordering","value":"ORDER_UNORDERED"},{"key":"packet_connection","value":"connection-0"}]}]}],"info":"","gas_wanted":"200000","gas_used":"54365","tx":null,"timestamp":""}
```
## 测试
测试加密功能正常
```
planetd tx blog send-ibc-post blog channel-0 "Hello" "Hello Mars, nihao" --from alice --chain-id earth --home ~/.earth
{"body":{"messages":[{"@type":"/cosmonaut.planet.blog.MsgSendIbcPost","sender":"cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d","port":"blog","channelID":"channel-0","timeoutTimestamp":"1631768377228065418","title":"Hello","content":"Hello Mars, nihao"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
{"height":"269","txhash":"876D4298D58C51EBBB69742BFD98F1F01E07648F727F9998FF60F798E3A346FC","codespace":"","code":0,"data":"0A0D0A0B53656E64496263506F7374","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"SendIbcPost\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]},{\"type\":\"send_packet\",\"attributes\":[{\"key\":\"packet_data\",\"value\":\"\\u0012P\\n\\u0005Hello\\u0012\\u0018SGVsbG8gTWFycywgbmloYW8=\\u001a-cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d\"},{\"key\":\"packet_timeout_height\",\"value\":\"0-0\"},{\"key\":\"packet_timeout_timestamp\",\"value\":\"1631768377228065418\"},{\"key\":\"packet_sequence\",\"value\":\"2\"},{\"key\":\"packet_src_port\",\"value\":\"blog\"},{\"key\":\"packet_src_channel\",\"value\":\"channel-0\"},{\"key\":\"packet_dst_port\",\"value\":\"blog\"},{\"key\":\"packet_dst_channel\",\"value\":\"channel-0\"},{\"key\":\"packet_channel_ordering\",\"value\":\"ORDER_UNORDERED\"},{\"key\":\"packet_connection\",\"value\":\"connection-0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"SendIbcPost"},{"key":"module","value":"ibc_channel"}]},{"type":"send_packet","attributes":[{"key":"packet_data","value":"\u0012P\n\u0005Hello\u0012\u0018SGVsbG8gTWFycywgbmloYW8=\u001a-cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d"},{"key":"packet_timeout_height","value":"0-0"},{"key":"packet_timeout_timestamp","value":"1631768377228065418"},{"key":"packet_sequence","value":"2"},{"key":"packet_src_port","value":"blog"},{"key":"packet_src_channel","value":"channel-0"},{"key":"packet_dst_port","value":"blog"},{"key":"packet_dst_channel","value":"channel-0"},{"key":"packet_channel_ordering","value":"ORDER_UNORDERED"},{"key":"packet_connection","value":"connection-0"}]}]}],"info":"","gas_wanted":"200000","gas_used":"54215","tx":null,"timestamp":""}
```
接收链可以正常接受并解密传输字段
```
planetd q blog list-post --node tcp://localhost:26659
Post:
- content: Hello Mars, I'm Alice from Earth
  creator: blog-channel-0-cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d
  id: "0"
  title: Hello
- content: Hello Mars, nihao
  creator: blog-channel-0-cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d
  id: "1"
  title: Hello
pagination:
  next_key: null
  total: "0"
```
传输链可以显示传输信息
```
planetd q blog list-sent-post
SentPost:
- chain: blog-channel-0
  creator: cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d
  id: "0"
  postID: "0"
  title: Hello
- chain: blog-channel-0
  creator: cosmos1f6la7c2rdpprgyqfz543pdqwyuv6lleezz7r7d
  id: "1"
  postID: "1"
  title: Hello
pagination:
  next_key: null
  total: "0"
```
# 以下是原文档
## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/cosmonaut/planet@latest! | sudo bash
```
`cosmonaut/planet` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
