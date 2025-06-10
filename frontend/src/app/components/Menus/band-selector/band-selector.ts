import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { StationConfigService } from '../../../services/station-config';
import { EnumBand } from '../../../../types';

@Component({
  selector: 'app-band-selector',
  imports: [FormsModule],
  templateUrl: './band-selector.html',
  styleUrl: './band-selector.css'
})
export class BandSelector {
  constructor(
    public StationConfig: StationConfigService
  ){
    // Initialize the selected band from the service
    this.selectedBand = this.StationConfig.band;
  }

  public bands = Object.values(EnumBand);
  public selectedBand: EnumBand = EnumBand.M80; // Default band

  onBandChange(): void {
    this.StationConfig.setStationConfig(
      this.StationConfig.callsign,
      this.StationConfig.name,
      this.selectedBand,
      this.StationConfig.mode
    )
    console.log(`Band changed to: ${this.selectedBand}`);
  }
}
