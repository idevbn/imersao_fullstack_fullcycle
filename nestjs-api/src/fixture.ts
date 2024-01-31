import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { getDataSourceToken } from '@nestjs/typeorm';
import { DataSource } from 'typeorm';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.init();

  const dataSource = app.get<DataSource>(getDataSourceToken());
  await dataSource.synchronize(true);

  const productRepo = dataSource.getRepository('Product');
  await productRepo.insert([
    {
      id: '8c2bdf10-2b1b-46c5-b5d5-a09e94759040',
      name: 'Product 1',
      description: 'Product 1 description',
      price: 100,
      image_url: 'http://localhost/9000/products/1.png',
    },
    {
      id: '4f423e4e-9272-4ba5-9e91-26cbe91d0cb6',
      name: 'Product 2',
      description: 'Product 2 description',
      price: 200,
      image_url: 'http://localhost/9000/products/2.png',
    },
    {
      id: '748c7229-6a53-4b02-bb38-4160f49c3c3e',
      name: 'Product 3',
      description: 'Product 3 description',
      price: 300,
      image_url: 'http://localhost/9000/products/3.png',
    },
    {
      id: '661f82ac-9bb1-4709-b3e2-f1a197a27562',
      name: 'Product 4',
      description: 'Product 4 description',
      price: 400,
      image_url: 'http://localhost/9000/products/4.png',
    },
    {
      id: 'c71fea5e-9e33-46e1-a9ae-6800175cbcc7',
      name: 'Product 5',
      description: 'Product 5 description',
      price: 500,
      image_url: 'http://localhost/9000/products/5.png',
    },
    {
      id: 'bbb5a152-86f2-448a-a88c-364642b7d601',
      name: 'Product 6',
      description: 'Product 6 description',
      price: 600,
      image_url: 'http://localhost/9000/products/6.png',
    },
    {
      id: '25e46c8f-1247-4575-bf02-7873b96dea79',
      name: 'Product 7',
      description: 'Product 7 description',
      price: 700,
      image_url: 'http://localhost/9000/products/7.png',
    },
    {
      id: '4d3658c5-55d3-47f2-a196-0071ac8cb4da',
      name: 'Product 8',
      description: 'Product 8 description',
      price: 800,
      image_url: 'http://localhost/9000/products/8.png',
    },
    {
      id: '85cd9afd-9705-4e07-8986-cc4da02bf3a6',
      name: 'Product 9',
      description: 'Product 9 description',
      price: 900,
      image_url: 'http://localhost/9000/products/9.png',
    },
    {
      id: 'd2816f63-f085-4080-aee5-b0eba4e46058',
      name: 'Product 10',
      description: 'Product 10 description',
      price: 1000,
      image_url: 'http://localhost/9000/products/10.png',
    },
  ]);

  await app.close();
}
bootstrap();
