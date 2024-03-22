FROM golang:latest

WORKDIR /app

# Copia os arquivos go.mod e go.sum para o diretório de trabalho dentro do contêiner
COPY ../go.mod .
COPY ../go.sum .

# Baixa as dependências do módulo
RUN go mod download

# Copie o restante do código-fonte para dentro do contêiner
COPY docker .

# Compile o aplicativo Go
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
