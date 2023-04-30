from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    # host = "https://e-commerce-api.fly.dev/api/v1"
    host = "127.0.0.1:8080/api/v1"

    @task
    def get_home(self):
        self.client.get("")

    @task
    def get_users(self):
        self.client.get("/users")

    @task
    def get_products(self):
        self.client.get("/products")