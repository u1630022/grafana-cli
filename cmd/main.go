package main

import (
    "github.com/grafana/grafana-openapi-client-go/client"
    "github.com/go-openapi/strfmt"
    "crypto/tls"
    "os"
    "log"
    "fmt"
    "strings"
)

const (
    grafana_secret_path = "./GRAFANA_KEY"
)

func main() {
    logger := log.Default()

    data, err := os.ReadFile(grafana_secret_path)
    if err != nil {
        logger.Fatalf("Reading secret from file produced err %s", err.Error())
    }
    key := strings.TrimSpace(string(data))

    c := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
        Host: "localhost:3000",
        BasePath: "/api",
        Schemes: []string{"http"},
        APIKey: key,
        TLSConfig: &tls.Config{},
        Debug: false,
    })

    dash, err := c.Dashboards.GetDashboardByUID("eds5h4owpwcg0b")
    if err != nil {
        logger.Fatalf("Unable to fetch dashboard: %s", err.Error())
    }

    fmt.Printf("Received dashboard '%s'\n", dash.Payload.Meta.Slug)
}
