"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const app = (0, express_1.default)();
app.use(express_1.default.json());
app.get('/', (_, res) => {
    const color = process.env.COLOR || 'unknown';
    res.json({ color });
});
app.get('/health', (_, res) => {
    res.json({ status: 'OK' });
});
app.get('/pid', (_, res) => {
    res.json({ pid: process.pid });
});
const port = process.env.PORT || 3000;
app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
