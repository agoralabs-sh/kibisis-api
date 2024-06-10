import 'swagger-ui/dist/swagger-ui.css';
import React from 'react';
import { createRoot, type Root } from 'react-dom/client';
import SwaggerUI from 'swagger-ui-react';

export async function onDOMContentLoaded(): Promise<void> {
  const _functionName = 'onDOMContentLoaded';
  const rootElement = document.getElementById('root');
  let root: Root;

  if (!rootElement) {
    console.error(
      `${_functionName}: failed to find "root" element to render react`
    );

    return;
  }

  root = createRoot(rootElement);

  root.render(
    <SwaggerUI url={__SWAGGER_SPEC_URL__} />
  );
}

window.addEventListener('DOMContentLoaded', onDOMContentLoaded);
