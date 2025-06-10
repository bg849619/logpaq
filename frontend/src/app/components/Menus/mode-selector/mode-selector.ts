import { Component } from '@angular/core';
import { StationConfigService } from '../../../services/station-config';
import { EnumMode } from '../../../../types';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-mode-selector',
  imports: [FormsModule],
  templateUrl: './mode-selector.html',
  styleUrl: './mode-selector.css'
})
export class ModeSelector {
  public selectedMode: EnumMode = EnumMode.SSB; // Default mode
  public modes = Object.values(EnumMode);
  
  constructor(
    public StationConfig: StationConfigService
  ) {
    this.selectedMode = this.StationConfig.mode; // Initialize the selected mode from the service
  }

  onModeChange(): void {
    this.StationConfig.setStationConfig(
      this.StationConfig.callsign,
      this.StationConfig.name,
      this.StationConfig.band,
      this.selectedMode
    );
    console.log(`Mode changed to: ${this.selectedMode}`);
  }
}
