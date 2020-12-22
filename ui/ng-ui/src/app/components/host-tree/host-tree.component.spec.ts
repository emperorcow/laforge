import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HostTreeComponent } from './host-tree.component';

describe('HostTreeComponent', () => {
  let component: HostTreeComponent;
  let fixture: ComponentFixture<HostTreeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HostTreeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HostTreeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
