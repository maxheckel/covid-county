import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CountyStatsComponent } from './county-stats.component';

describe('CountyStatsComponent', () => {
  let component: CountyStatsComponent;
  let fixture: ComponentFixture<CountyStatsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CountyStatsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CountyStatsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
