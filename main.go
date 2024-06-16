package main

import (
    "fmt"
    "strings"
)

// Processor struct and predefined values
type Processor struct {
    Manufacturer string
    Brand        string
    Modifier     string
    Generation   int
    Cores        int
    ClockSpeed   float64
}

var processorManufacturers = []string{"Intel", "AMD"}
var intelBrands = []string{"Core", "Pentium", "Celeron"}
var amdBrands = []string{"Ryzen", "Athlon"}
var intelModifiers = []string{"i3", "i5", "i7", "i9"}
var amdModifiers = []string{"3", "5", "7", "9"}
var intelGenerations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
var amdGenerations = []int{1, 2, 3, 4, 5, 6}
var processorCores = []int{1, 2, 4, 6, 8, 10, 12, 16}

// GPU struct and predefined values
type GPU struct {
    Form         string
    Manufacturer string
    BrandPrefix  string
    Cores        int
    VRAM         int
}

var gpuForms = []string{"Integrated", "Dedicated"}
var integratedManufacturers = []string{"Intel", "AMD"}
var dedicatedManufacturers = []string{"Intel", "NVIDIA", "AMD"}
var intelIntegratedBrandPrefixes = []string{"HD", "UHD", "Iris Plus", "Iris Xe"}
var amdIntegratedBrandPrefixes = []string{"Radeon R", "Vega", "HD"}
var nvidiaGPUs = []string{"GeForce RTX", "GeForce GTX", "GeForce GT", "GeForce MX"}
var intelDedicatedGPUs = []string{"Iris Xe"}
var amdDedicatedGPUs = []string{"Radeon RX", "Vega", "Pro"}

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

var predefinedActivities = []string{
    "Programming and Development",
    "Photography",
    "Video Editing and Production",
    "Designing",
    "Business and Work",
    "Academics",
    "Audio Production",
    "Architecture and 3D Design",
    "Animations and 3D Rendering",
    "Church Presentation",
    "Data Science and Analytics",
    "Machine Learning",
    "Gaming",
    "Entertainment",
}

func recommendActivities(laptop Laptop) []string {
    var recommendedActivities []string

    // Analyze processor
    switch {
    case laptop.Processor.Cores >= 8 && laptop.Processor.ClockSpeed >= 3.5:
        recommendedActivities = append(recommendedActivities, "Video Editing and Production", "Animations and 3D Rendering", "Data Science and Analytics", "Machine Learning")
    case laptop.Processor.Cores >= 6 && laptop.Processor.ClockSpeed >= 3.0:
        recommendedActivities = append(recommendedActivities, "Video Editing and Production", "Animations and 3D Rendering", "Architecture and 3D Design")
    case laptop.Processor.Cores >= 4 && laptop.Processor.ClockSpeed >= 2.5:
        recommendedActivities = append(recommendedActivities, "Programming and Development", "Business and Work", "Academics")
    default:
        recommendedActivities = append(recommendedActivities, "Business and Work", "Academics", "Entertainment")
    }

    // Analyze GPU
    if laptop.GPU.Form == "Dedicated" {
        if laptop.GPU.VRAM >= 8 {
            recommendedActivities = append(recommendedActivities, "Gaming", "Animations and 3D Rendering", "Machine Learning")
        } else if laptop.GPU.VRAM >= 4 {
            recommendedActivities = append(recommendedActivities, "Gaming", "Animations and 3D Rendering")
        }
    } else if laptop.GPU.Form == "Integrated" {
        recommendedActivities = append(recommendedActivities, "Photography", "Designing", "Church Presentation")
    }

    // Analyze RAM
    if laptop.RAM.SizeInGB >= 32 {
        recommendedActivities = append(recommendedActivities, "Data Science and Analytics", "Machine Learning")
    } else if laptop.RAM.SizeInGB >= 16 {
        recommendedActivities = append(recommendedActivities, "Programming and Development", "Business and Work", "Video Editing and Production")
    } else if laptop.RAM.SizeInGB >= 8 {
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

func main() {
    // Prompt user for input
    fmt.Println("Enter laptop specifications:")

    var processor Processor
    var gpu GPU
    var ram RAM
    var storage Storage

    // Processor input
    fmt.Print("Processor Manufacturer: ")
    var manufacturerInput string
    fmt.Scanln(&manufacturerInput)
    processor.Manufacturer = strings.ToLower(manufacturerInput)

    fmt.Print("Processor Brand: ")
    var brandInput string
    fmt.Scanln(&brandInput)
    processor.Brand = strings.ToLower(brandInput)

    fmt.Print("Processor Modifier: ")
    var modifierInput string
    fmt.Scanln(&modifierInput)
    processor.Modifier = strings.ToLower(modifierInput)

    fmt.Print("Processor Generation: ")
    fmt.Scanln(&processor.Generation)

    fmt.Print("Processor Cores: ")
    fmt.Scanln(&processor.Cores)

    fmt.Print("Processor Clock Speed (GHz): ")
    fmt.Scanln(&processor.ClockSpeed)

    // GPU input
    fmt.Print("GPU Form (Integrated/Dedicated): ")
    var formInput string
    fmt.Scanln(&formInput)
    gpu.Form = strings.ToLower(formInput)

    fmt.Print("GPU Manufacturer: ")
    var gpuManufacturerInput string
    fmt.Scanln(&gpuManufacturerInput)
    gpu.Manufacturer = strings.ToLower(gpuManufacturerInput)

    fmt.Print("GPU Brand & Prefix: ")
    var brandPrefixInput string
    fmt.Scanln(&brandPrefixInput)
    gpu.BrandPrefix = strings.ToLower(brandPrefixInput)

    if gpu.Form == "dedicated" {
        fmt.Print("GPU Cores: ")
        fmt.Scanln(&gpu.Cores)

        fmt.Print("GPU VRAM (GB): ")
        fmt.Scanln(&gpu.VRAM)
    }

    // RAM input
    fmt.Print("RAM Size (GB): ")
    fmt.Scanln(&ram.SizeInGB)

    // Storage input
    fmt.Print("Storage Size (GB): ")
    fmt.Scanln(&storage.SizeInGB)
    fmt.Print("Storage Type (SSD/HDD): ")
    var storageType string
    fmt.Scanln(&storageType)
    storage.TypeSSD = strings.ToLower(storageType) == "ssd"

    // Create Laptop struct
    laptop := Laptop{
        Processor: processor,
        GPU:       gpu,
        RAM:       ram,
        Storage:   storage,
    }

    // Recommend activities
    recommendedActivities := recommendActivities(laptop)
    fmt.Println("\nRecommended Activities:")
    for _, activity := range recommendedActivities {
        fmt.Println("-", activity)
    }
}