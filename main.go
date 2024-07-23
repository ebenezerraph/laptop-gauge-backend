package main

import (
	"encoding/json"
	"net/http"
)

// Structs
type Req struct {
	Manufacturer string  `json:"manufacturer"`
	Brand        string  `json:"brand"`
	Modifier     string  `json:"modifier"`
	Generation   int     `json:"generation"`
	Cores        int     `json:"cores"`
	ClockSpeed   float64 `json:"clockSpeed"`
}

type Processor struct {
	Manufacturer string
	Brand        string
	Modifier     string
	Generation   int
	Cores        int
	ClockSpeed   float64
}

type GPU struct {
	Form         string
	Manufacturer string
	BrandPrefix  string
	Cores        int
	VRAM         int
}

type RAM struct {
	SizeInGB int
}

type Storage struct {
	SizeInGB int
	TypeSSD  bool
}

type Laptop struct {
	Processor Processor
	GPU       GPU
	RAM       RAM
	Storage   Storage
}

// Constants
const (
	TierHighEnd    = "high-end"
	TierHighMid    = "high-mid"
	TierLowMid     = "low-mid"
	TierEntryLevel = "entry-level"
)

// Processor analysis
func getProcessorTier(p Processor) string {
	coresClockGenTier := analyzeCoresClockGeneration(p.Cores, p.ClockSpeed, p.Generation)

	switch p.Manufacturer {
	case "intel":
		return getIntelTier(p.Brand, p.Modifier, coresClockGenTier)
	case "amd":
		return getAMDTier(p.Brand, p.Modifier, coresClockGenTier)
	case "apple":
		return getAppleTier(p.Brand)
	default:
		return TierEntryLevel
	}
}

func analyzeCoresClockGeneration(cores int, clockSpeed float64, generation int) string {
	switch {
	case cores >= 8 && clockSpeed >= 3.5 && generation >= 10:
		return TierHighEnd
	case cores >= 6 && clockSpeed >= 3.0 && generation >= 8:
		return TierHighMid
	case cores >= 4 && clockSpeed >= 2.5 && generation >= 6:
		return TierLowMid
	default:
		return TierEntryLevel
	}
}

func getIntelTier(brand, modifier, coresClockGenTier string) string {
	if brand == "core" {
		switch modifier {
		case "i9":
			return getHigherTier(coresClockGenTier, TierHighMid)
		case "i7":
			return getLowerTier(coresClockGenTier, TierHighMid)
		case "i5":
			return getLowerTier(coresClockGenTier, TierLowMid)
		default:
			return TierEntryLevel
		}
	}
	return TierEntryLevel
}

func getAMDTier(brand, modifier, coresClockGenTier string) string {
	if brand == "ryzen" {
		switch modifier {
		case "9":
			return getHigherTier(coresClockGenTier, TierHighMid)
		case "7":
			return getLowerTier(coresClockGenTier, TierHighMid)
		case "5":
			return getLowerTier(coresClockGenTier, TierLowMid)
		default:
			return TierEntryLevel
		}
	}
	return TierEntryLevel
}

func getAppleTier(brand string) string {
	switch brand {
	case "m3":
		return TierHighEnd
	case "m2":
		return TierHighMid
	case "m1":
		return TierLowMid
	default:
		return TierEntryLevel
	}
}

func getHigherTier(currentTier, defaultTier string) string {
	if currentTier == TierHighEnd {
		return TierHighEnd
	}
	return defaultTier
}

func getLowerTier(currentTier, defaultTier string) string {
	if currentTier == defaultTier {
		return defaultTier
	}
	return getNextLowerTier(defaultTier)
}

func getNextLowerTier(tier string) string {
	switch tier {
	case TierHighEnd:
		return TierHighMid
	case TierHighMid:
		return TierLowMid
	default:
		return TierEntryLevel
	}
}

// GPU analysis
func getGPUTier(gpu GPU) string {
	if gpu.Form == "dedicated" {
		return getDedicatedGPUTier(gpu)
	}
	return getIntegratedGPUTier(gpu)
}

func getDedicatedGPUTier(gpu GPU) string {
	vramCoresTier := analyzeVRAMCores(gpu.VRAM, gpu.Cores)

	switch gpu.Manufacturer {
	case "nvidia":
		return getNvidiaTier(gpu.BrandPrefix, vramCoresTier)
	case "amd":
		return getAMDGPUTier(gpu.BrandPrefix, vramCoresTier)
	case "intel":
		if gpu.BrandPrefix == "iris xe" {
			return TierHighMid
		}
	}
	return TierEntryLevel
}

func getIntegratedGPUTier(gpu GPU) string {
	switch gpu.Manufacturer {
	case "intel":
		return getIntelIntegratedTier(gpu.BrandPrefix)
	case "amd":
		return getAMDIntegratedTier(gpu.BrandPrefix)
	}
	return TierEntryLevel
}

func analyzeVRAMCores(vram, cores int) string {
	switch {
	case vram >= 8 && cores >= 4096:
		return TierHighEnd
	case vram >= 4 && cores >= 2048:
		return TierHighMid
	case vram >= 2 && cores >= 1024:
		return TierLowMid
	default:
		return TierEntryLevel
	}
}

func getNvidiaTier(brandPrefix, vramCoresTier string) string {
	switch brandPrefix {
	case "geforce rtx":
		return getHigherTier(vramCoresTier, TierHighMid)
	case "geforce gtx":
		return getLowerTier(vramCoresTier, TierHighMid)
	case "geforce gt":
		return getLowerTier(vramCoresTier, TierLowMid)
	default:
		return TierEntryLevel
	}
}

func getAMDGPUTier(brandPrefix, vramCoresTier string) string {
	switch brandPrefix {
	case "radeon rx":
		return getHigherTier(vramCoresTier, TierHighMid)
	case "radeon vega":
		return getLowerTier(vramCoresTier, TierHighMid)
	case "radeon pro":
		return getLowerTier(vramCoresTier, TierLowMid)
	default:
		return TierEntryLevel
	}
}

func getIntelIntegratedTier(brandPrefix string) string {
	switch brandPrefix {
	case "iris xe":
		return TierHighMid
	case "iris plus":
		return TierLowMid
	default:
		return TierEntryLevel
	}
}

func getAMDIntegratedTier(brandPrefix string) string {
	switch brandPrefix {
	case "radeon vega":
		return TierHighMid
	case "radeon r":
		return TierLowMid
	default:
		return TierEntryLevel
	}
}

// Activity recommendation
func recommendActivities(laptop Laptop) []string {
	activities := []string{}

	processorTier := getProcessorTier(laptop.Processor)
	activities = append(activities, getProcessorActivities(processorTier)...)

	gpuTier := getGPUTier(laptop.GPU)
	activities = append(activities, getGPUActivities(gpuTier, laptop.GPU.Form)...)

	activities = append(activities, getRAMActivities(laptop.RAM.SizeInGB)...)

	if laptop.Storage.TypeSSD {
		activities = append(activities, "Video Editing and Production", "Audio Production")
	}

	return limitAndDeduplicate(activities, 5)
}

func getProcessorActivities(tier string) []string {
	switch tier {
	case TierHighEnd:
		return []string{"Video Editing and Production", "Animations and 3D Rendering", "Data Science and Analytics", "Machine Learning"}
	case TierHighMid:
		return []string{"Video Editing and Production", "Animations and 3D Rendering", "Architecture and 3D Design"}
	case TierLowMid:
		return []string{"Programming and Development", "Business and Work", "Academics"}
	default:
		return []string{"Business and Work", "Academics", "Entertainment"}
	}
}

func getGPUActivities(tier, form string) []string {
	switch tier {
	case TierHighEnd:
		return []string{"Gaming", "Animations and 3D Rendering", "Machine Learning"}
	case TierHighMid:
		return []string{"Gaming", "Animations and 3D Rendering"}
	case TierLowMid:
		if form == "Integrated" {
			return []string{"Photography", "Designing", "Church Presentation"}
		}
		return []string{"Gaming"}
	default:
		if form == "Integrated" {
			return []string{"Photography", "Designing", "Church Presentation"}
		}
	}
	return []string{}
}

func getRAMActivities(sizeInGB int) []string {
	switch {
	case sizeInGB >= 16:
		return []string{"Data Science and Analytics", "Machine Learning"}
	case sizeInGB >= 8:
		return []string{"Programming and Development", "Business and Work", "Video Editing and Production"}
	case sizeInGB >= 4:
		return []string{"Programming and Development", "Business and Work"}
	}
	return []string{}
}

func limitAndDeduplicate(activities []string, limit int) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, activity := range activities {
		if !seen[activity] {
			seen[activity] = true
			result = append(result, activity)
			if len(result) == limit {
				break
			}
		}
	}
	return result
}

// HTTP handler (assuming this is part of a web service)
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the request and send response
	// (Implementation details omitted for brevity)
}

func main() {
	// Set up HTTP server and routes
	// (Implementation details omitted for brevity)
}
