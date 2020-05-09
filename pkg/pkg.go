package pkg

const XAPIKey = "x-apikey"

type LastAnalysisStats struct {
	HarmLess   int
	Undetected int
	Suspicious int
	Malicious  int
}
