export enum EnumBand {
    M160 = "160m",
    M80 = "80m",
    M40 = "40m",
    M20 = "20m",
    M15 = "15m",
    M10 = "10m",
    M6 = "6m",
    M2 = "2m",
    M125 = "1.25m",
    M70 = "70cm",
}

export enum EnumMode {
    CW = "CW",
    SSB = "SSB",
    FM = "FM",
    DIGI = "DIGI",
    AM = "AM",
}

export interface LogConfig {
    callsign: string;
    class: string;
    section: string;
}

export interface StationConfig {
    callsign: string;
    name: string;
    band: EnumBand;
    mode: EnumMode;
}