import { Type } from 'class-transformer';
import {
  ArrayNotEmpty,
  IsInt,
  IsNotEmpty,
  IsPositive,
  IsString,
  MaxLength,
  ValidateNested,
} from 'class-validator';

export class CreateOrderDto {
  @ArrayNotEmpty()
  @Type(() => OrderItemDto)
  @ValidateNested({ each: true })
  items: OrderItemDto[];

  @IsString()
  @IsNotEmpty()
  @MaxLength(255)
  card_hash: string;
}

export class OrderItemDto {
  @IsInt()
  @IsPositive()
  @IsNotEmpty()
  quantity: number;

  @IsString()
  @IsNotEmpty()
  @MaxLength(36)
  product_id: string;
}
