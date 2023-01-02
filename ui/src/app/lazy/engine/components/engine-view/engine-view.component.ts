import { Component, Input } from '@angular/core';
import { EngineInfo } from 'src/app/static/models';

@Component({
  selector: 'app-engine-view',
  templateUrl: './engine-view.component.html',
  styleUrls: ['./engine-view.component.css']
})
export class EngineViewComponent {

  @Input() engines!: EngineInfo[] | null;
  constructor() {}
}
