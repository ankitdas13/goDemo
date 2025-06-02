# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#NotesUnionParam">NotesUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#NotesUnion">NotesUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderListPaymentsResponse">OrderListPaymentsResponse</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}/payments">client.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderListPaymentsParams">OrderListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#OrderListPaymentsResponse">OrderListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#AcquirerData">AcquirerData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentGetParams">PaymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentUpdateParams">PaymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/capture">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.Capture">Capture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentCaptureParams">PaymentCaptureParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/refund">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.NewRefund">NewRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentNewRefundParams">PaymentNewRefundParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.ListRefunds">ListRefunds</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentListRefundsParams">PaymentListRefundsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/card">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.GetCard">GetCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds/{refund_id}">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentService.GetRefund">GetRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentGetRefundParams">PaymentGetRefundParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## QrCodes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#QrCode">QrCode</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeListPaymentsResponse">PaymentQrCodeListPaymentsResponse</a>

Methods:

- <code title="post /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeNewParams">PaymentQrCodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeGetParams">PaymentQrCodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeListParams">PaymentQrCodeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/qr_codes/{id}/close">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}/payments">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeListPaymentsParams">PaymentQrCodeListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentQrCodeListPaymentsResponse">PaymentQrCodeListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundList">RefundList</a>

# PaymentLinks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLink">PaymentLink</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkListResponse">PaymentLinkListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>

Methods:

- <code title="post /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkNewParams">PaymentLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkUpdateParams">PaymentLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkListParams">PaymentLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkListResponse">PaymentLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment_links/{id}/notify_by/{medium}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkService.Notify">Notify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, medium <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkNotifyParamsMedium">PaymentLinkNotifyParamsMedium</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkNotifyParams">PaymentLinkNotifyParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Methods:

- <code title="get /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#RefundUpdateNotesParams">RefundUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settlements

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Settlement">Settlement</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementListResponse">SettlementListResponse</a>

Methods:

- <code title="get /settlements/{id}">client.Settlements.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Settlement">Settlement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements">client.Settlements.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementListParams">SettlementListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementListResponse">SettlementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recon

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementReconGetResponse">SettlementReconGetResponse</a>

Methods:

- <code title="get /settlements/recon/combined">client.Settlements.Recon.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementReconService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementReconGetParams">SettlementReconGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementReconGetResponse">SettlementReconGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ondemand

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemand">SettlementOndemand</a>

Methods:

- <code title="post /settlements/ondemand">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemandNewParams">SettlementOndemandNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements/ondemand/{id}">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemandGetParams">SettlementOndemandGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Payout">Payout</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts/{id}">client.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Payout">Payout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Plans

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PlanItemParam">PlanItemParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Plan">Plan</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PlanItem">PlanItem</a>

Methods:

- <code title="post /plans">client.Plans.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PlanService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PlanNewParams">PlanNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Plan">Plan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /plans/{id}">client.Plans.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#PlanService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#Plan">Plan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#InvoiceNewResponse">InvoiceNewResponse</a>

Methods:

- <code title="post /invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#InvoiceNewParams">InvoiceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go">godemo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/godemo-go#InvoiceNewResponse">InvoiceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
