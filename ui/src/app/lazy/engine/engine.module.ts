import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { RouterModule, Routes } from '@angular/router';
import { StoreModule } from '@ngrx/store';
import * as fromEngineState from './state';
import * as fromEngineReducer from './state/engine.reducer';
import { EngineViewComponent } from './components/engine-view/engine-view.component';
import { EngineListComponent } from './components/engine-list/engine-list.component';
import { EffectsModule } from '@ngrx/effects';
import { EngineEffects } from './state/engine.effect';

const routes: Routes = [{ path: '', component: EngineListComponent }];

const EFFECTS = [EngineEffects];

@NgModule({
  declarations: [EngineViewComponent, EngineListComponent],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    HttpClientModule,
    StoreModule.forFeature(fromEngineState.engineFeatureKey, fromEngineReducer.engineInfoReducer),
    EffectsModule.forFeature(EFFECTS),
  ],
})
export class EngineModule {}
