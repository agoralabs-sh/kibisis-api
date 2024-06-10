// types
import type IResponse from './IResponse';

type TGlobal = typeof globalThis & {
  main: () => Promise<IResponse>;
};

export default TGlobal;
