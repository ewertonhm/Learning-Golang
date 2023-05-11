package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	ServerName    string
	ServerURL     string
	TempoExecucao int64
	Status        int
	DataFalha     string
}

func (s *Server) checkServer() {
	now := time.Now()
	get, err := http.Get(s.ServerURL)

	if err != nil {
		fmt.Printf("Server %s is down [%s]\n", s.ServerName, err.Error())
		s.Status = 0
		s.DataFalha = now.Format("02/01/2006 15:04:05")
		return
	}

	s.Status, s.TempoExecucao = get.StatusCode, time.Since(now).Milliseconds()

	if s.Status != 200 {
		fmt.Printf("Service: [%s] - Status: [%d] Tempo de carga: [%dms]\n", s.ServerName, s.Status, s.TempoExecucao)
		s.DataFalha = now.Format("02/01/2006 15:04:05")
	} else {
		fmt.Printf("Service: [%s] - Status: [%d] Tempo de carga: [%dms]\n", s.ServerName, s.Status, s.TempoExecucao)
	}
}

func checkServers(servidores []Server) []Server {
	var downServers []Server

	for _, servidor := range servidores {
		servidor.checkServer()
		if servidor.Status != 200 {
			downServers = append(downServers, servidor)
		}

	}

	return downServers
}

func criarListaServidores(serverList *os.File) []Server {
	csvReader := csv.NewReader(serverList)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var servidores []Server
	for i, line := range data {
		if i > 0 {
			servidor := Server{
				ServerName: line[0],
				ServerURL:  line[1],
			}
			servidores = append(servidores, servidor)
		}
	}
	return servidores
}

func openFiles(serverListFile string, downtimeFile string) (*os.File, *os.File) {
	ServerList, err := os.OpenFile(serverListFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	downtimeList, err := os.OpenFile(downtimeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return ServerList, downtimeList
}

func addDowntimeToFile(f *os.File, s []Server) {
	csvWriter := *csv.NewWriter(f)
	for _, servidor := range s {
		line := []string{servidor.ServerName, servidor.ServerURL, servidor.DataFalha, fmt.Sprintf("%d", servidor.Status), fmt.Sprintf("%d", servidor.TempoExecucao)}
		csvWriter.Write(line)
	}
	csvWriter.Flush()
}

func main() {
	serverList, downtimeList := openFiles(os.Args[1], os.Args[2])
	servidores := criarListaServidores(serverList)

	downServers := checkServers(servidores)
	addDowntimeToFile(downtimeList, downServers)

	defer serverList.Close()
	defer downtimeList.Close()
}
