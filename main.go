package main

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