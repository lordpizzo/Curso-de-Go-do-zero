package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

// Grar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	// Inicializa a aplicação
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Flags:  flags,
			Usage:  "Busca IPs na Internet",
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Flags:  flags,
			Usage:  "Busca Servidores na Internet",
			Action: buscaServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	println("Buscando IPs para", host)

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		println(ip.String())
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")
	println("Buscando Servidores para", host)

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		println(servidor.Host)
	}
}
