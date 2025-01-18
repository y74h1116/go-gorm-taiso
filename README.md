# GORM 朝の体操
- GORM（Go言語の ORM）のトレーニングメニューです。  
  - PostgreSQL と Go言語の実行環境を含めてあります。  
  - PostgreSQL には users テーブルが作られます。  
  - Go言語の実行環境には、GORM がインストールされます。  
  - GORM で、users テーブルを操作する練習ができます。  
- Docker はインストールして起動しておいてください。
- 準備：Github からソースを取得したら以下を実行  
    ```
    → Mac ならターミナルで、Windows なら WSL で実行

      ※ソースを配置したディレクトリへ移動
    $ cd go-gorm-taiso
      ※イメージを生成
    $ docker compose build  
      ※コンテナを起動
    $ docker compose up  

    → 以下は別タブまたはウィンドウを開いて実行

      ※docker ps の結果、go-gorm-taiso と db-taiso が Up になっていればOK
    $ docker ps  
      ※GORM など必要なものを取得してインストール
    $ docker exec -it go-gorm-taiso go mod tidy  


      ※Go 言語のプログラムの実行
    $ docker exec -it go-gorm-taiso go run main.go
    ```
- main.go ファイルの下の方に練習問題が書いてありますので、そこに実装を打ち込んで練習できます。  
- example.go に解答例を記載しました。  
- users テーブルの定義は、[init.sql](./storage/postgres/init/init.sql) に記載してあります。  
