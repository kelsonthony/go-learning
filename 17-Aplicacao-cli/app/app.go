package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação CLI"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "https://example.com",
		},
	}

	// Definir as flags e comandos aqui, se necessário
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs e Nomes de Servidores",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca Nomes de Servidores",
			Flags:  flags,
			Action: buscarNomesServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	fmt.Printf("Buscando IPs e Nomes de Servidores em: %s\n", host)

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Printf("IP encontrado: %s\n", ip.String())
	}
}

func buscarNomesServidores(c *cli.Context) {
	host := c.String("host")
	fmt.Printf("Buscando Nomes de Servidores em: %s\n", host)

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Printf("Servidor encontrado: %s\n", servidor.Host)
	}
}
