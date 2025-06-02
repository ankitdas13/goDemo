# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#NotesUnionParam">NotesUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#NotesUnion">NotesUnion</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderListPaymentsResponse">OrderListPaymentsResponse</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}/payments">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderListPaymentsParams">OrderListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#OrderListPaymentsResponse">OrderListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#AcquirerData">AcquirerData</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentGetParams">PaymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentUpdateParams">PaymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/capture">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.Capture">Capture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentCaptureParams">PaymentCaptureParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/refund">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.NewRefund">NewRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentNewRefundParams">PaymentNewRefundParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.ListRefunds">ListRefunds</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentListRefundsParams">PaymentListRefundsParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/card">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.GetCard">GetCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds/{refund_id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentService.GetRefund">GetRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentGetRefundParams">PaymentGetRefundParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## QrCodes

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#QrCode">QrCode</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeListPaymentsResponse">PaymentQrCodeListPaymentsResponse</a>

Methods:

- <code title="post /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeNewParams">PaymentQrCodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeGetParams">PaymentQrCodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeListParams">PaymentQrCodeListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/qr_codes/{id}/close">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}/payments">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeListPaymentsParams">PaymentQrCodeListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentQrCodeListPaymentsResponse">PaymentQrCodeListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundList">RefundList</a>

# PaymentLinks

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLink">PaymentLink</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkListResponse">PaymentLinkListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>

Methods:

- <code title="post /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkNewParams">PaymentLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkUpdateParams">PaymentLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkListParams">PaymentLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkListResponse">PaymentLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment_links/{id}/notify_by/{medium}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkService.Notify">Notify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, medium <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkNotifyParamsMedium">PaymentLinkNotifyParamsMedium</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkNotifyParams">PaymentLinkNotifyParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Methods:

- <code title="get /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#RefundUpdateNotesParams">RefundUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settlements

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Settlement">Settlement</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementListResponse">SettlementListResponse</a>

Methods:

- <code title="get /settlements/{id}">client.Settlements.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Settlement">Settlement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements">client.Settlements.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementListParams">SettlementListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementListResponse">SettlementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recon

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementReconGetResponse">SettlementReconGetResponse</a>

Methods:

- <code title="get /settlements/recon/combined">client.Settlements.Recon.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementReconService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementReconGetParams">SettlementReconGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementReconGetResponse">SettlementReconGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ondemand

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemand">SettlementOndemand</a>

Methods:

- <code title="post /settlements/ondemand">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemandNewParams">SettlementOndemandNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements/ondemand/{id}">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemandGetParams">SettlementOndemandGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Payout">Payout</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts/{id}">client.Payouts.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Payout">Payout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Plans

Params Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PlanItemParam">PlanItemParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Plan">Plan</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PlanItem">PlanItem</a>

Methods:

- <code title="post /plans">client.Plans.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PlanService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PlanNewParams">PlanNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Plan">Plan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /plans/{id}">client.Plans.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#PlanService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#Plan">Plan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#InvoiceNewResponse">InvoiceNewResponse</a>

Methods:

- <code title="post /invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#InvoiceNewParams">InvoiceNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo">godemo</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goDemo#InvoiceNewResponse">InvoiceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
