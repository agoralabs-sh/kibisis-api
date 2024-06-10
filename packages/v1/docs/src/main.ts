import { readFile } from 'fs/promises';
import { resolve } from 'path';

// types
import type { IResponse, TGlobal } from '@app/types';

export async function main(): Promise<IResponse> {
  const html = await readFile(resolve(__dirname, 'index.html'));

  return {
    headers: {
      'Content-Type': 'text/html'
    },
    statusCode: 200,
    body: html.toString(),
  };
}

(global as TGlobal).main = main;
