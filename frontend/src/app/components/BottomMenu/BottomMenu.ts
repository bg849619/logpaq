import { Component } from '@angular/core';
import { LogConfigService } from '../../services/log-config';
import { StationConfigService } from '../../services/station-config';
import { MatDialog } from '@angular/material/dialog';
import { LogConfigMenu } from '../Menus/log-config-menu/log-config-menu';
import { OperatorMenu } from '../Menus/operator-menu/operator-menu';
import { BandSelector } from '../Menus/band-selector/band-selector';
import { ModeSelector } from '../Menus/mode-selector/mode-selector';

@Component({
  selector: 'app-bottom-menu',
  imports: [BandSelector, ModeSelector],
  templateUrl: './BottomMenu.html',
  styleUrl: './BottomMenu.css',
  providers: [ LogConfigService, StationConfigService ]
})
export class BottomMenu {
  constructor(
    public LogConfig: LogConfigService,
    public StationConfig: StationConfigService,
    public logConfigDialog: MatDialog,
    public operatorDialog: MatDialog
  ) { 
  }

  ngOnInit(): void {
    this.LogConfig.getLogConfig()
    this.StationConfig.getStationConfig();
  }

  openLogConfigMenu(): void {
    const dialogRef = this.logConfigDialog.open(LogConfigMenu, {
      width: '250px',
      data: {...this.LogConfig}
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed with result:', result);
      if (result) {
        this.LogConfig.setLogConfig(result.callsign, result.class, result.section);
      }
    });
  }

  openOperatorDialog(): void {
    const dialogRef = this.operatorDialog.open(OperatorMenu, {
      width: '250px',
      data: {...this.StationConfig }// Do this to get a clone string.
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The operator dialog was closed with result:', result);
      if (result) {
        this.StationConfig.setStationConfig(result.callsign, result.name, this.StationConfig.band, this.StationConfig.mode);
      }
    });
  }
}