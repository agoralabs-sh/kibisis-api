import 'swagger-ui/dist/swagger-ui.css';
import React from 'react';
import { createRoot, type Root } from 'react-dom/client';
import SwaggerUI from 'swagger-ui-react';

// spec
import spec from '@app/spec/swagger.json';

export async function onDOMContentLoaded(): Promise<void> {
  const _functionName: string = 'onDOMContentLoaded';
  const rootElement: Element | null = document.getElementById('root');
  let root: Root;

  if (!rootElement) {
    console.error(
      `${_functionName}: failed to find "root" element to render react`
    );

    return;
  }

  root = createRoot(rootElement);

  root.render(
    <SwaggerUI spec={spec} />
  );
}

window.addEventListener('DOMContentLoaded', onDOMContentLoaded);
