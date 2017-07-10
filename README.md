# qiita2esa-cli

## これは何？
Qiita::Teamからesaに記事を移行するcliです。

### やること
- 記事の移行
- 記事の投稿者引き継ぎ
    - esaのscreen_nameとqiitaのidが異なる場合はesa_botとして投稿します。
- Qiitaの記事名に`/`がついている場合は`-`に置換します

### やらないこと
- コメントの移行
- 画像の移行

## Getting Start
```bash
$ go get github.com/sawadashota/qiita2esa-cli
```

## Usage
```bash
$ qiita2esa-cli -q [Qiita::Team name] -qToken [qiita token] -eToken [esa token] [-e [esa team name] -restart-from [Process ID]]
```

|Flag|説明|備考|
|:--|:--|:--|
|-q|Qiita::Teamのチーム名|必須|
|-qToken|Qiitaのアクセストークン|必須|
|-eToken|esaのアクセストークン|必須|
|-e|esaチーム名|Qiita::Teamのチーム名と異なるときのみ必須|
|-restart-from|何個目の記事から処理を再開させるか|任意。数値のみ|