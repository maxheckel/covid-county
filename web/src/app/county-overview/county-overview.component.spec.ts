import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CountyOverviewComponent } from './county-overview.component';

describe('CountyOverviewComponent', () => {
  let component: CountyOverviewComponent;
  let fixture: ComponentFixture<CountyOverviewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CountyOverviewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CountyOverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
