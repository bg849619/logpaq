import { TestBed } from '@angular/core/testing';

import { LogConfig } from './log-config';

describe('LogConfig', () => {
  let service: LogConfig;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LogConfig);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
