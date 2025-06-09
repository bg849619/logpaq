import { Injectable } from '@angular/core';
import { Observable, of} from 'rxjs';
import { LogConfig } from '../../types';

@Injectable({
  providedIn: 'root'
})
export class LogConfigService {
  callsign: string = 'W8BAP'; // PLACEHOLDER
  class: string = '5A'; // PLACEHOLDER
  section: string = 'OH'; // PLACEHOLDER

  constructor( ) { }

  getLogConfig(): Observable<LogConfig>{
    const logConfig = of({
      callsign: this.callsign,
      class: this.class,
      section: this.section
    });
    return logConfig;
  }

  async setLogConfig(callsign: string, classType: string, section: string) {
    console.log(`Setting log config: ${callsign}, ${classType}, ${section}`);
    this.callsign = callsign;
    this.class = classType;
    this.section = section;
  }
}
