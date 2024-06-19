package main

import (
	"encoding/json"
	"log"
	"net/http"
    "os"
)

type Req struct {
    Manufacturer string  `json:"manufacturer"`
    Brand        string  `json:"brand"`
    Modifier     string  `json:"modifier"`
    Generation   int     `json:"generation"`
    Cores        int     `json:"cores"`
    ClockSpeed   float64 `json:"clockSpeed"`
}
err := json.NewDecoder(r.Body).Decode(&req)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
// Processor struct and predefined values
type Processor struct {
    Manufacturer string
    Brand        string
    Modifier     string
    Generation   int
    Cores        int
    ClockSpeed   float64
}

// GPU struct and predefined values
type GPU struct {
    Form         string
    Manufacturer string
    BrandPrefix  string
    Cores        int
    VRAM         int
}


// RAM struct
type RAM struct {
    SizeInGB int
}

// Storage struct
type Storage struct {
    SizeInGB int
    TypeSSD  bool
}

// Laptop struct
type Laptop struct {
    Processor Processor
    GPU       GPU
    RAM       RAM
    Storage   Storage
}

func analyzeCoresClockGeneration(cores int, clockSpeed float64, generation int) string {
    var tier string

    switch {
    case cores >= 8 && clockSpeed >= 3.5 && generation >= 10:
        tier = "high-end"
    case cores >= 6 && clockSpeed >= 3.0 && generation >= 8:
        tier = "high-mid"
    case cores >= 4 && clockSpeed >= 2.5 && generation >= 6:
        tier = "low-mid"
    default:
        tier = "entry-level"
    }

    return tier
}

func getProcessorTier(processor Processor) string {
    var tier string

    coresClockGenTier := analyzeCoresClockGeneration(processor.Cores, processor.ClockSpeed, processor.Generation)

    switch processor.Manufacturer {
    case "intel":
        switch processor.Brand {
        case "core":
            switch processor.Modifier {
            case "i9":
                if coresClockGenTier == "high-end" {
                    tier = "high-end"
                } else {
                    tier = "high-mid"
                }
            case "i7":
                if coresClockGenTier == "high-mid" {
                    tier = "high-mid"
                } else {
                    tier = "low-mid"
                }
            case "i5":
                if coresClockGenTier == "low-mid" {
                    tier = "low-mid"
                } else {
                    tier = "entry-level"
                }
            case "i3":
                tier = "entry-level"
            }
        case "pentium", "celeron":
            tier = "entry-level"
        }
    case "amd":
        switch processor.Brand {
        case "ryzen":
            switch processor.Modifier {
            case "9":
                if coresClockGenTier == "high-end" {
                    tier = "high-end"
                } else {
                    tier = "high-mid"
                }
            case "7":
                if coresClockGenTier == "high-mid" {
                    tier = "high-mid"
                } else {
                    tier = "low-mid"
                }
            case "5":
                if coresClockGenTier == "low-mid" {
                    tier = "low-mid"
                } else {
                    tier = "entry-level"
                }
            case "3":
                tier = "entry-level"
            }
        case "athlon":
            tier = "entry-level"
        }
    case "apple":
        switch processor.Brand {
        case "m3":
            tier = "high-end"
        case "m2":
            tier = "high-mid"
        case "m1":
            tier = "low-mid"
        }
    }

    return tier
}

func analyzeVRAMCores(vram int, cores int) string {
    var tier string

    switch {
    case vram >= 8 && cores >= 4096:
        tier = "high-end"
    case vram >= 4 && cores >= 2048:
        tier = "high-mid"
    case vram >= 2 && cores >= 1024:
        tier = "low-mid"
    default:
        tier = "entry-level"
    }

    return tier
}

func getGPUTier(gpu GPU) string {
    var tier string

    if gpu.Form == "dedicated" {
        vramCoresTier := analyzeVRAMCores(gpu.VRAM, gpu.Cores)

        switch gpu.Manufacturer {
        case "nvidia":
            switch gpu.BrandPrefix {
            case "geforce rtx":
                if vramCoresTier == "high-end" {
                    tier = "high-end"
                } else {
                    tier = "high-mid"
                }
            case "geforce gtx":
                if vramCoresTier == "high-mid" {
                    tier = "high-mid"
                } else {
                    tier = "low-mid"
                }
            case "geforce gt":
                if vramCoresTier == "low-mid" {
                    tier = "low-mid"
                } else {
                    tier = "entry-level"
                }
            case "geforce mx":
                tier = "entry-level"
            }
        case "intel":
            if gpu.BrandPrefix == "iris xe" {
                tier = "high-mid"
            }
        case "amd":
            switch gpu.BrandPrefix {
            case "radeon rx":
                if vramCoresTier == "high-end" {
                    tier = "high-end"
                } else {
                    tier = "high-mid"
                }
            case "radeon vega":
                if vramCoresTier == "high-mid" {
                    tier = "high-mid"
                } else {
                    tier = "low-mid"
                }
            case "radeon pro":
                if vramCoresTier == "low-mid" {
                    tier = "low-mid"
                } else {
                    tier = "entry-level"
                }
            }
        }
    } else if gpu.Form == "integrated" {
        switch gpu.Manufacturer {
        case "intel":
            switch gpu.BrandPrefix {
            case "iris xe":
                tier = "high-mid"
            case "iris plus":
                tier = "low-mid"
            case "uhd":
                tier = "entry-level"
            case "hd":
                tier = "entry-level"
            }
        case "amd":
            switch gpu.BrandPrefix {
            case "radeon vega":
                tier = "high-mid"
            case "radeon r":
                tier = "low-mid"
            case "radeon hd":
                tier = "entry-level"
            }
        }
    }

    return tier
}

func recommendActivities(laptop Laptop) []string {
    var recommendedActivities []string

    // Analyze processor
    processorTier := getProcessorTier(laptop.Processor)
    switch processorTier {
    case "high-end":
        recommendedActivities = append(recommendedActivities, "Video Editing and Production", "Animations and 3D Rendering", "Data Science and Analytics", "Machine Learning")
    case "high-mid":
        recommendedActivities = append(recommendedActivities, "Video Editing and Production", "Animations and 3D Rendering", "Architecture and 3D Design")
    case "low-mid":
        recommendedActivities = append(recommendedActivities, "Programming and Development", "Business and Work", "Academics")
    case "entry-level":
        recommendedActivities = append(recommendedActivities, "Business and Work", "Academics", "Entertainment")
    }

    // Analyze GPU
    gpuTier := getGPUTier(laptop.GPU)
    switch gpuTier {
    case "high-end":
        recommendedActivities = append(recommendedActivities, "Gaming", "Animations and 3D Rendering", "Machine Learning")
    case "high-mid":
        recommendedActivities = append(recommendedActivities, "Gaming", "Animations and 3D Rendering")
    case "low-mid":
        if laptop.GPU.Form == "Integrated" {
            recommendedActivities = append(recommendedActivities, "Photography", "Designing", "Church Presentation")
        } else {
            recommendedActivities = append(recommendedActivities, "Gaming")
        }
    case "entry-level":
        if laptop.GPU.Form == "Integrated" {
            recommendedActivities = append(recommendedActivities, "Photography", "Designing", "Church Presentation")
        }
    }

    // Analyze RAM
    if laptop.RAM.SizeInGB >= 16 {
        recommendedActivities = append(recommendedActivities, "Data Science and Analytics", "Machine Learning")
    } else if laptop.RAM.SizeInGB >= 8 {
        recommendedActivities = append(recommendedActivities, "Programming and Development", "Business and Work", "Video Editing and Production")
    } else if laptop.RAM.SizeInGB >= 4 {
        recommendedActivities = append(recommendedActivities, "Programming and Development", "Business and Work")
    }

    // Analyze Storage
    if laptop.Storage.TypeSSD {
        recommendedActivities = append(recommendedActivities, "Video Editing and Production", "Audio Production")
    }

    // Remove duplicates and limit recommended activities to 5
    recommendedActivities = removeDuplicates(recommendedActivities)
    if len(recommendedActivities) > 5 {
        recommendedActivities = recommendedActivities[:5]
    }

    return recommendedActivities
}

func removeDuplicates(activities []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range activities {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}