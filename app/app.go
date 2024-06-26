package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPS de enderecos na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "https://github.com/feliipecardosoo",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
