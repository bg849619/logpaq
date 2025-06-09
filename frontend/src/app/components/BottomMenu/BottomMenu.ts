import { Component } from '@angular/core';
import { LogConfigService } from '../../services/log-config';
import { StationConfigService } from '../../services/station-config';
import { EnumBand, EnumMode, LogConfig, StationConfig } from '../../../types';
import { MatDialog } from '@angular/material/dialog';
import { LogConfigMenu } from '../Menus/log-config-menu/log-config-menu';

@Component({
  selector: 'app-bottom-menu',
  imports: [],
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
}