import express from 'express';

const app = express();
const PORT = process.env.PORT || 5011;

app.get('/health', (req, res) => {
    res.json({ status: 'healthy', service: '11-email-service' });
});

app.listen(PORT, () => {
    console.log(`11-email-service running on port ${PORT}`);
});
