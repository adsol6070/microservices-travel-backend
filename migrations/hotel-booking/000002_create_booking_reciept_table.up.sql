CREATE TYPE booking_status AS ENUM (
    'pending',
    'confirmed',
    'awaiting_payment',
    'payment_received',
    'cancelled',
    'cancelled_by_guest',
    'cancelled_by_hotel',
    'checked_in',
    'checked_out',
    'no_show',
    'awaiting_confirmation',
    'pending_review',
    'reviewed',
    'expired',
    'in_progress',
    'under_payment_review',
    'under_cancellation_review',
    'refunded',
    'disputed'
);

CREATE TABLE hotel_bookings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_id VARCHAR(255) NOT NULL, 
    booking_status booking_status NOT NULL,
    confirmation_number VARCHAR(255) NOT NULL,
    check_in_date DATE NOT NULL,
    check_out_date DATE NOT NULL,
    guest_count INT NOT NULL,
    room_type VARCHAR(255) NOT NULL,
    room_quantity INT NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    tax_amount DECIMAL(10,2) NOT NULL,
    payment_type VARCHAR(255) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_card_vendor VARCHAR(50),
    payment_card_number VARCHAR(20),
    payment_card_expiry VARCHAR(10),
    payment_card_holder VARCHAR(255),
    hotel_id VARCHAR(255) NOT NULL,
    hotel_name VARCHAR(255) NOT NULL,
    hotel_chain_code VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE guests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id UUID REFERENCES hotel_bookings(id) ON DELETE CASCADE,
    guest_reference INT NOT NULL,
    title VARCHAR(10),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE associated_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id UUID REFERENCES hotel_bookings(id) ON DELETE CASCADE,
    reference VARCHAR(255) NOT NULL,
    origin_system_code VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
