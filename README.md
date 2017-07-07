# qiita2esa-cli

## これは何？
Qiita::Teamからesaに記事を移行するcliです。

### やること
- 記事の移行
- 記事の投稿者引き継ぎ
    - esaのscreen_nameとqiitaのidが異なる場合はesa_botとして投稿します。

### やらないこと
- コメントの移行
- 画像の移行

## Getting Start
```bash
$ go get github.com/sawadashota/qiita2esa-cli
```

## Usage
```bash
$ qiita2esa-cli -q [Qiita::Team name] -qToken [qiita token] -eToken [esa token] [-e [esa team name]]
```

|Flag|説明|備考|
|:--|:--|:--|
|-q|Qiita::Teamのチーム名|必須|
|-qToken|Qiitaのアクセストークン|必須|
|-eToken|esaのアクセストークン|必須|
|-e|esaチーム名|Qiita::Teamのチーム名と異なるときのみ必須|