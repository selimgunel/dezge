import { ErrorEffects } from './error.effects';
import { SnackBarEffects } from './snack-bar.effects';
import { SpinnerEffects } from './spinner.effects';
import { Web3GatewayEffects } from './web3-gateway.effects';

export const effects: any[] = [ErrorEffects, SnackBarEffects, SpinnerEffects,
    Web3GatewayEffects];

export * from './error.effects';
export * from './snack-bar.effects';
export * from './spinner.effects';
export * from './web3-gateway.effects';

