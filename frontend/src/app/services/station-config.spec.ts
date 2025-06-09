import { TestBed } from '@angular/core/testing';

import { StationConfig } from './station-config';

describe('StationConfig', () => {
  let service: StationConfig;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(StationConfig);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
