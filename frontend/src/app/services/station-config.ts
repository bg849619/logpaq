import { Injectable } from '@angular/core';
import { EnumBand, EnumMode } from '../../types';

@Injectable({
  providedIn: 'root'
})
export class StationConfigService {
  callsign: string = 'KE8PAQ'; // PLACEHOLDER
  name: string = 'Test Station'; // PLACEHOLDER
  band: EnumBand = EnumBand.M80; // Default band
  mode: EnumMode = EnumMode.SSB; // Default mode

  constructor() { }

  getStationConfig() {
    return {
      callsign: this.callsign,
      name: this.name,
      band: this.band,
      mode: this.mode
    };
  }

  setStationConfig(callsign: string, name: string, band: EnumBand, mode: EnumMode) {
    this.callsign = callsign;
    this.name = name;
    this.band = band;
    this.mode = mode;
  }
}
