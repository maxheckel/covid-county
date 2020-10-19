import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DailyCountsComponent } from './daily-counts.component';

describe('DailyCountsComponent', () => {
  let component: DailyCountsComponent;
  let fixture: ComponentFixture<DailyCountsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DailyCountsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DailyCountsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
