# ターミナルからクリップボードに画像をコピーしたい

インストール

```
    go install github.com/tf63/go-copy-img@latest
```

### goenv (ローカルで使う場合の環境)

```
    brew install goenv

    # bashの場合
    export GOENV_ROOT=$HOME/.goenv
    export PATH=$GOENV_ROOT/bin:$PATH

    # fishの場合
    set -Ux GOENV_ROOT $HOME/.goenv
    fish_add_path $GOENV_ROOT/bin
```

インストール可能なGoのバージョンを確認

```
    goenv install -l
```

インストール

```
    goenv install 1.19

    # インストールされているGoのバージョンを確認
    goenv versions
```

ローカルのバージョン指定

```
    goenv local 1.19

    # 確認
    go version
```

https://zenn.dev/erueru_tech/articles/13e8512d9312de
