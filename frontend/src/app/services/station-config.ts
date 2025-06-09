import { Injectable } from '@angular/core';
import { EnumBand, EnumMode } from '../../types';

@Injectable({
  providedIn: 'root'
})
export class StationConfigService {
  operator: string = 'KE8PAQ'; // PLACEHOLDER
  band: EnumBand = EnumBand.M80; // Default band
  mode: EnumMode = EnumMode.SSB; // Default mode

  constructor() { }

  getStationConfig() {
    return {
      operator: this.operator,
      band: this.band,
      mode: this.mode
    };
  }

  setStationConfig(operator: string, band: EnumBand, mode: EnumMode) {
    this.operator = operator;
    this.band = band;
    this.mode = mode;
  }
}
