import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LogConfigMenu } from './log-config-menu';

describe('LogConfigMenu', () => {
  let component: LogConfigMenu;
  let fixture: ComponentFixture<LogConfigMenu>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LogConfigMenu]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LogConfigMenu);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
