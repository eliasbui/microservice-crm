import express, { Request, Response } from 'express';
import { createClient } from 'redis';
import amqp from 'amqplib';
import winston from 'winston';
import dotenv from 'dotenv';

dotenv.config();

// Logger setup
const logger = winston.createLogger({
  level: 'info',
  format: winston.format.json(),
  transports: [
    new winston.transports.Console(),
  ],
});

const app = express();
app.use(express.json());

// Configuration
const PORT = process.env.NOTIFICATION_SERVICE_PORT || 8084;
const RABBITMQ_URL = `amqp://${process.env.RABBITMQ_USER}:${process.env.RABBITMQ_PASSWORD}@${process.env.RABBITMQ_HOST}:${process.env.RABBITMQ_PORT}${process.env.RABBITMQ_VHOST}`;

// Types
interface EmailNotification {
  to: string;
  subject: string;
  body: string;
  from?: string;
}

interface SMSNotification {
  to: string;
  message: string;
}

interface PushNotification {
  userId: string;
  title: string;
  body: string;
  data?: Record<string, any>;
}

// Health check endpoints
app.get('/health', (_req: Request, res: Response) => {
  res.json({ status: 'healthy', service: 'notification-service' });
});

app.get('/ready', (_req: Request, res: Response) => {
  res.json({ status: 'ready', service: 'notification-service' });
});

app.get('/live', (_req: Request, res: Response) => {
  res.json({ status: 'alive', service: 'notification-service' });
});

// API endpoints
app.post('/api/notifications/email', async (req: Request, res: Response) => {
  try {
    const notification: EmailNotification = req.body;
    logger.info('Sending email notification', { to: notification.to });

    // TODO: Implement actual email sending with SendGrid/SMTP
    await sendEmail(notification);

    res.json({
      success: true,
      message: 'Email notification sent',
      notificationId: generateId()
    });
  } catch (error) {
    logger.error('Failed to send email', { error });
    res.status(500).json({ success: false, error: 'Failed to send email' });
  }
});

app.post('/api/notifications/sms', async (req: Request, res: Response) => {
  try {
    const notification: SMSNotification = req.body;
    logger.info('Sending SMS notification', { to: notification.to });

    // TODO: Implement actual SMS sending
    await sendSMS(notification);

    res.json({
      success: true,
      message: 'SMS notification sent',
      notificationId: generateId()
    });
  } catch (error) {
    logger.error('Failed to send SMS', { error });
    res.status(500).json({ success: false, error: 'Failed to send SMS' });
  }
});

app.post('/api/notifications/push', async (req: Request, res: Response) => {
  try {
    const notification: PushNotification = req.body;
    logger.info('Sending push notification', { userId: notification.userId });

    // TODO: Implement actual push notification sending
    await sendPushNotification(notification);

    res.json({
      success: true,
      message: 'Push notification sent',
      notificationId: generateId()
    });
  } catch (error) {
    logger.error('Failed to send push notification', { error });
    res.status(500).json({ success: false, error: 'Failed to send push notification' });
  }
});

// Service implementations
async function sendEmail(notification: EmailNotification): Promise<void> {
  // Placeholder implementation
  logger.info('Email sent successfully', { to: notification.to });
}

async function sendSMS(notification: SMSNotification): Promise<void> {
  // Placeholder implementation
  logger.info('SMS sent successfully', { to: notification.to });
}

async function sendPushNotification(notification: PushNotification): Promise<void> {
  // Placeholder implementation
  logger.info('Push notification sent successfully', { userId: notification.userId });
}

function generateId(): string {
  return `notif_${Date.now()}_${Math.random().toString(36).substring(7)}`;
}

// RabbitMQ consumer setup
async function setupMessageConsumer() {
  try {
    const connection = await amqp.connect(RABBITMQ_URL);
    const channel = await connection.createChannel();

    const queue = 'notifications';
    await channel.assertQueue(queue, { durable: true });

    logger.info('RabbitMQ consumer connected, waiting for messages...');

    channel.consume(queue, async (msg) => {
      if (msg) {
        const content = JSON.parse(msg.content.toString());
        logger.info('Received notification message', { type: content.type });

        // Process based on notification type
        switch (content.type) {
          case 'email':
            await sendEmail(content.data);
            break;
          case 'sms':
            await sendSMS(content.data);
            break;
          case 'push':
            await sendPushNotification(content.data);
            break;
        }

        channel.ack(msg);
      }
    });
  } catch (error) {
    logger.error('Failed to setup message consumer', { error });
    // Retry after delay
    setTimeout(setupMessageConsumer, 5000);
  }
}

// Start server
app.listen(PORT, () => {
  logger.info(`Notification Service started on port ${PORT}`);
  setupMessageConsumer().catch(console.error);
});
