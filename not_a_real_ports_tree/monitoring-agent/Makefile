# $FreeBSD$

PORTNAME=	monitoring_agent
PORTVERSION=	1
PORTREVISION=	1
CATEGORIES=	net-mgmt
MASTER_SITES=	# empty
DISTFILES=	# empty
EXTRACT_ONLY=	# empty

MAINTAINER=	infrastructure@codeweavers.net
COMMENT=	Monitoring Agent

LICENSE=	APACHE20

RUN_DEPENDS=	

NO_BUILD=	yes
NO_MTREE=	yes

SUB_FILES=	pkg-install pkg-deinstall
SUB_LIST=	PORTNAME=${PORTNAME}

do-install:
		${MKDIR} ${WRKSRC}
		${MKDIR} ${STAGEDIR}${PREFIX}/bin
		cp files/usr/local/bin/monitoring-agent ${WRKSRC}/monitoring_agent
		${INSTALL_PROGRAM} ${WRKSRC}/monitoring_agent ${STAGEDIR}${PREFIX}/bin/monitoring_agent
		#${INSTALL} ${WRKSRC}/monitoring_agent ${STAGEDIR}${PREFIX}/bin/monitoring_agent

.include <bsd.port.mk>

