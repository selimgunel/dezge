import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { EngineInfo } from 'src/app/static/models';

import * as fromEngine from './../../state';
import * as engineActions from '../../state/engine.action';

@Component({
  selector: 'app-engine-list',
  template: `
    <div class="grid place-items-center v-screen">
      <app-engine-view [engines]="engines$ | async"> </app-engine-view>
    </div>
  `,
  styleUrls: ['./engine-list.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class EngineListComponent implements OnInit {
  engines$: Observable<EngineInfo[] | null>;
  constructor(private store: Store) {
    this.engines$ = store.select(fromEngine.getEngineList);
  }
  ngOnInit() {
    this.store.dispatch(engineActions.list());
  }
}
