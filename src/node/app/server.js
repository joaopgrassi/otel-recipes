"use strict";

const api = require('@opentelemetry/api');
const tracer = require('./tracer')("nodejs-recipe");

const express = require("express");
const axios = require("axios");
const app = express();

const PORT = process.env.PORT || "8080";

app.get("/", async(req, res) => {

    createManualSpan();

    let externalCall = await axios.get("https://inspiration.goprogram.ai/");
    return res.send(externalCall.data);
});

function createManualSpan() {
    const parentSpan = api.trace.getSpan(api.context.active());
    const ctx = api.trace.setSpan(api.context.active(), parentSpan);
    const span = tracer.startSpan('manual-span', undefined, ctx);

    span.setAttribute('attribute_key', 'attribute_value');

    span.addEvent('invoking createManualSpan');

    span.end();
}

app.listen(parseInt(PORT, 10), () => {
    console.log(`Listening for requests on http://localhost:${PORT}`);
});