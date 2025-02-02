import { serve } from "bun";
import Stripe from "stripe";

console.log(process.env.STRIPE_SECRET_KEY);
const stripe = new Stripe(process.env.STRIPE_SECRET_KEY as string);

serve({
  port: 3001,
  fetch: async (req) => {
    if (req.method === "POST" && new URL(req.url).pathname === "/create-payment-intent") {
      console.log("Received request to create payment intent");
      fetch("http://localhost:8000/orders/2/")
        .then(res => res.ok)
        .then(data => {
          console.log(data);
        })
        .catch(err => {
          console.log(err);
        });

      try {
        const paymentIntent = await stripe.paymentIntents.create({
          amount: 1000, // Amount in cents
          currency: "usd",
          automatic_payment_methods: {
            enabled: true
          }
        });

        return new Response(JSON.stringify({ clientSecret: paymentIntent.client_secret }), {
          headers: { "Content-Type": "application/json" },
          status: 200,
        });
      } catch (error) {
        return new Response(JSON.stringify({ error: error.message }), {
          headers: { "Content-Type": "application/json" },
          status: 500,
        });
      }
    }

    if (req.method === "GET" && new URL(req.url).pathname === "/success") {
      console.log("Received request to show success page");
    }

    // Handle invalid routes
    return new Response("Not Found", { status: 404 });
  },

});

console.log("Server running on http://localhost:3001");
