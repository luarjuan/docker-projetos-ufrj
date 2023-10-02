FROM golang:1.18-alpine
#usa o diretorio da pasta definida
WORKDIR /dockergitlab 
#usa os arquivos para a linguagem go que foram disponibilizados no gitlab
COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN go build -o api.go
#define a porta de entrada que foi solicitada
EXPOSE 3000

#executa a aplicação api.go
CMD ["go", "run", "api.go"]

