import express from 'express';

const app = express();
const PORT = process.env.PORT || 5021;

app.get('/health', (req, res) => {
    res.json({ status: 'healthy', service: '21-dashboard-service' });
});

app.listen(PORT, () => {
    console.log(`21-dashboard-service running on port ${PORT}`);
});
