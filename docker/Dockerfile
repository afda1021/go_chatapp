# どのイメージを使うか、追加でいろいろ設定
FROM golang:1.15.6-alpine3.12
# RUN→コマンドの実行
RUN mkdir /go/src/chatapp
# RUNとかCMDで定義したコマンドをどのディレクトリで実行するか
WORKDIR /go/src/chatapp

ENV GO111MODULE=on

# github.com/pilu/fresh：ファイルを監視し、更新されたらホットリロード
RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh
    
CMD ["fresh"]

# 8080番ポートを公開する
# EXPOSE 8080