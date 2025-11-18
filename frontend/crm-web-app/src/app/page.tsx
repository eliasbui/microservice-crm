'use client';

import { Box, Container, Typography, Button, Grid, Card, CardContent } from '@mui/material';
import {
  People,
  TrendingUp,
  Chat,
  Notifications,
  Analytics
} from '@mui/icons-material';
import Link from 'next/link';

export default function HomePage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 8 }}>
        {/* Hero Section */}
        <Box sx={{ textAlign: 'center', mb: 8 }}>
          <Typography
            variant="h2"
            component="h1"
            gutterBottom
            sx={{ fontWeight: 700, mb: 2 }}
          >
            Welcome to CRM System
          </Typography>
          <Typography
            variant="h5"
            color="text.secondary"
            paragraph
            sx={{ mb: 4 }}
          >
            Modern Customer Relationship Management with AI-Powered Intelligence
          </Typography>
          <Box sx={{ display: 'flex', gap: 2, justifyContent: 'center' }}>
            <Button
              variant="contained"
              size="large"
              component={Link}
              href="/login"
            >
              Get Started
            </Button>
            <Button
              variant="outlined"
              size="large"
              component={Link}
              href="/dashboard"
            >
              View Dashboard
            </Button>
          </Box>
        </Box>

        {/* Features Grid */}
        <Grid container spacing={4}>
          <Grid item xs={12} md={4}>
            <Card sx={{ height: '100%' }}>
              <CardContent sx={{ textAlign: 'center', p: 4 }}>
                <People sx={{ fontSize: 60, color: 'primary.main', mb: 2 }} />
                <Typography variant="h5" gutterBottom>
                  Customer Management
                </Typography>
                <Typography color="text.secondary">
                  Manage your customers, track interactions, and build stronger relationships
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={4}>
            <Card sx={{ height: '100%' }}>
              <CardContent sx={{ textAlign: 'center', p: 4 }}>
                <TrendingUp sx={{ fontSize: 60, color: 'primary.main', mb: 2 }} />
                <Typography variant="h5" gutterBottom>
                  Opportunity Tracking
                </Typography>
                <Typography color="text.secondary">
                  Track sales opportunities through your pipeline and close more deals
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={4}>
            <Card sx={{ height: '100%' }}>
              <CardContent sx={{ textAlign: 'center', p: 4 }}>
                <Chat sx={{ fontSize: 60, color: 'primary.main', mb: 2 }} />
                <Typography variant="h5" gutterBottom>
                  AI-Powered Chat
                </Typography>
                <Typography color="text.secondary">
                  Get intelligent assistance with AI-powered chat for customer insights
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={6}>
            <Card sx={{ height: '100%' }}>
              <CardContent sx={{ textAlign: 'center', p: 4 }}>
                <Notifications sx={{ fontSize: 60, color: 'primary.main', mb: 2 }} />
                <Typography variant="h5" gutterBottom>
                  Smart Notifications
                </Typography>
                <Typography color="text.secondary">
                  Stay informed with real-time notifications via email, SMS, and push
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={6}>
            <Card sx={{ height: '100%' }}>
              <CardContent sx={{ textAlign: 'center', p: 4 }}>
                <Analytics sx={{ fontSize: 60, color: 'primary.main', mb: 2 }} />
                <Typography variant="h5" gutterBottom>
                  Advanced Analytics
                </Typography>
                <Typography color="text.secondary">
                  Get insights from your data with advanced analytics and reporting
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        </Grid>

        {/* Stats Section */}
        <Box sx={{ mt: 8, textAlign: 'center' }}>
          <Grid container spacing={4}>
            <Grid item xs={12} md={3}>
              <Typography variant="h3" color="primary.main" gutterBottom>
                10K+
              </Typography>
              <Typography variant="h6" color="text.secondary">
                Active Users
              </Typography>
            </Grid>
            <Grid item xs={12} md={3}>
              <Typography variant="h3" color="primary.main" gutterBottom>
                50K+
              </Typography>
              <Typography variant="h6" color="text.secondary">
                Customers Managed
              </Typography>
            </Grid>
            <Grid item xs={12} md={3}>
              <Typography variant="h3" color="primary.main" gutterBottom>
                95%
              </Typography>
              <Typography variant="h6" color="text.secondary">
                Satisfaction Rate
              </Typography>
            </Grid>
            <Grid item xs={12} md={3}>
              <Typography variant="h3" color="primary.main" gutterBottom>
                24/7
              </Typography>
              <Typography variant="h6" color="text.secondary">
                AI Support
              </Typography>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Container>
  );
}
