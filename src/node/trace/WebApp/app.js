'use strict';

const express = require('express')
const app = express()

//span_dependencies_start
const api = require ('@opentelemetry/api');
//span_dependencies_end

app.get('/', (req, res) => {
  //span_creation_start
  const tracer = api.trace.getTracer("RecipeTracer", "0.1.0");
  const span = tracer.startSpan("HelloWorldSpan", { attributes: { attribute1 : 'value1' } });
  span.end();
  //span_creation_end

  res.send('Hello World!')
})

app.listen(8080, () => console.log('Server listening on http://localhost:8080'))