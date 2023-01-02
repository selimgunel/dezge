import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EngineViewComponent } from './engine-view.component';

describe('EngineViewComponent', () => {
  let component: EngineViewComponent;
  let fixture: ComponentFixture<EngineViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EngineViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EngineViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
